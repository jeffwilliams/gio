// Code generated by build.go. DO NOT EDIT.

package headless

import "gioui.org/gpu/backend"

var (
	shader_input_vert = backend.ShaderSources{
		Inputs:    []backend.InputLocation{backend.InputLocation{Name: "position", Location: 0, Semantic: "POSITION", SemanticIndex: 0, Type: 0x0, Size: 4}},
		GLSL100ES: "\nattribute vec4 position;\n\nvoid main()\n{\n    gl_Position = position;\n}\n\n",
		GLSL300ES: "#version 300 es\n\nlayout(location = 0) in vec4 position;\n\nvoid main()\n{\n    gl_Position = position;\n}\n\n",
		GLSL130:   "#version 130\n\nin vec4 position;\n\nvoid main()\n{\n    gl_Position = position;\n}\n\n",
		/*
		   static float4 gl_Position;
		   static float4 position;

		   struct SPIRV_Cross_Input
		   {
		       float4 position : POSITION;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float4 gl_Position : SV_Position;
		   };

		   void vert_main()
		   {
		       gl_Position = position;
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       position = stage_input.position;
		       vert_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.gl_Position = gl_Position;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
	shader_simple_frag = backend.ShaderSources{
		GLSL100ES: "precision mediump float;\nprecision highp int;\n\nvoid main()\n{\n    gl_FragData[0] = vec4(0.25, 0.550000011920928955078125, 0.75, 1.0);\n}\n\n",
		GLSL300ES: "#version 300 es\nprecision mediump float;\nprecision highp int;\n\nlayout(location = 0) out vec4 fragColor;\n\nvoid main()\n{\n    fragColor = vec4(0.25, 0.550000011920928955078125, 0.75, 1.0);\n}\n\n",
		GLSL130:   "#version 130\n\nout vec4 fragColor;\n\nvoid main()\n{\n    fragColor = vec4(0.25, 0.550000011920928955078125, 0.75, 1.0);\n}\n\n",
		/*
		   static float4 fragColor;

		   struct SPIRV_Cross_Output
		   {
		       float4 fragColor : SV_Target0;
		   };

		   void frag_main()
		   {
		       fragColor = float4(0.25f, 0.550000011920928955078125f, 0.75f, 1.0f);
		   }

		   SPIRV_Cross_Output main()
		   {
		       frag_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.fragColor = fragColor;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
	shader_simple_vert = backend.ShaderSources{
		GLSL100ES: "\nvoid main()\n{\n    float x;\n    float y;\n    if (gl_VertexID == 0)\n    {\n        x = 0.0;\n        y = 0.5;\n    }\n    else\n    {\n        if (gl_VertexID == 1)\n        {\n            x = 0.5;\n            y = -0.5;\n        }\n        else\n        {\n            x = -0.5;\n            y = -0.5;\n        }\n    }\n    gl_Position = vec4(x, y, 0.5, 1.0);\n}\n\n",
		GLSL300ES: "#version 300 es\n\nvoid main()\n{\n    float x;\n    float y;\n    if (gl_VertexID == 0)\n    {\n        x = 0.0;\n        y = 0.5;\n    }\n    else\n    {\n        if (gl_VertexID == 1)\n        {\n            x = 0.5;\n            y = -0.5;\n        }\n        else\n        {\n            x = -0.5;\n            y = -0.5;\n        }\n    }\n    gl_Position = vec4(x, y, 0.5, 1.0);\n}\n\n",
		GLSL130:   "#version 130\n\nvoid main()\n{\n    float x;\n    float y;\n    if (gl_VertexID == 0)\n    {\n        x = 0.0;\n        y = 0.5;\n    }\n    else\n    {\n        if (gl_VertexID == 1)\n        {\n            x = 0.5;\n            y = -0.5;\n        }\n        else\n        {\n            x = -0.5;\n            y = -0.5;\n        }\n    }\n    gl_Position = vec4(x, y, 0.5, 1.0);\n}\n\n",
		/*
		   static float4 gl_Position;
		   static int gl_VertexIndex;
		   struct SPIRV_Cross_Input
		   {
		       uint gl_VertexIndex : SV_VertexID;
		   };

		   struct SPIRV_Cross_Output
		   {
		       float4 gl_Position : SV_Position;
		   };

		   void vert_main()
		   {
		       float x;
		       float y;
		       if (gl_VertexIndex == 0)
		       {
		           x = 0.0f;
		           y = 0.5f;
		       }
		       else
		       {
		           if (gl_VertexIndex == 1)
		           {
		               x = 0.5f;
		               y = -0.5f;
		           }
		           else
		           {
		               x = -0.5f;
		               y = -0.5f;
		           }
		       }
		       gl_Position = float4(x, y, 0.5f, 1.0f);
		   }

		   SPIRV_Cross_Output main(SPIRV_Cross_Input stage_input)
		   {
		       gl_VertexIndex = int(stage_input.gl_VertexIndex);
		       vert_main();
		       SPIRV_Cross_Output stage_output;
		       stage_output.gl_Position = gl_Position;
		       return stage_output;
		   }

		*/
		HLSL: []byte(nil),
	}
)